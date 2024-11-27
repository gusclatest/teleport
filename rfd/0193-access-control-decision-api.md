---
authors: Alan Parra (alan.parra@goteleport.com), Forrest Marshall (forrest@goteleport.com)
state: draft
---

# RFD 0193 - Access Control Decision API

## Required approvers

- Engineering: @rosstimothy

## What

This proposal describes a means of reworking teleport access-control decisions to remove the need to perform
primary access-control decisions on peripheral agents and to bring teleport's internals more in line with common
access-control best practices.

## Why

Access-control decisions in teleport are typically made at the agent, and often times a single operation
will incur many small access-control decisions over the course of its execution. This distributed and innervated
model of access-control decision making has some benefits. For example, early versions of teleport made it a
priority to have agents continue to be able to make access-control decisions even when the auth service was
unavailable for long periods of time. However, there are also serious drawbacks to the distributed and innervated
model:

- Scalability: Access-control decisions at the agent require that agents have access to all necessary information
in order to be able to make correct access-control decisions. This forces us to encode a large amount of otherwise
unnecessary information into certificates, and forces agents to cache all resources relevant to access-control
decision making (namely roles). In some larger clusters, role replication alone can account for a significant portion
of all load incurred after an interruption of auth connectivity.

- Compatibility: Because teleport access-control includes the ability to specify unscoped deny rules, there is no
sane way to represent many new access-control features to older agents s.t. they can continue to make sensible
decisions. We've typically had to resort to cumbersome and imperfect hacks that negatively impact user experience,
such as injecting wildcard deny rules into the caches of outdated agents to force them to deny any access that might
require knowledge of a new feature.

- Maintainability: From both the developer and user perspective, distributed and innervated access-control decision
making poses maintainability challenges. For developers, it is difficult to reason about or change access-control
features because the implementation is so spread out, and so many different teleport components rely on internal
details of the implementation. For users, in order to take advantage of a bug fix or new feature, they must often
update thousands of teleport installations. This is particularly worrisome in the case of security fixes, as
access-control decision making is a particularly sensitive part of teleport.

- Development Cadence: Due mostly to the aforementioned issues, it is very difficult to make any non-trivial changes
to teleport access-control decision making. Most major proposed changes to teleport access-control decision making
(e.g. scoped rbac) have never even reached the point of a working prototype. This isn't only because of the lack of
a clear abstraction boundary, but that is a very significant contributing factor.

## Details

### Key Terminology

- Policy Decision Point (PDP): A component responsible for making access-control decisions. Sometimes referred to as
an Access Decision Function in some literature.

- Policy Enforcement Point (PEP): A component responsible for enforcing access-control decisions. Sometimes referred to
as an Access Enforcement Function in some literature.

- Decision Request: A request sent to the PDP describing an action and requesting a decision on whether or not
the action should be allowed, and any necessary information needed to correctly enforce limitations upon the action.

- Permit: A single data structure encoding all necessary information for a PEP to correctly enforce a conditional/
parameterized allow decision.

### Overview

We will rework teleport's access-control decision making to have a well-defined boundary between Policy Decision Point
and Policy Enforcement Point. The goal of this rework will be to abstract as much complexity as possible into the PDP,
while attempting to minimize the need to actually change how or where enforcement is done today. The PDP/PEP boundary
will be representable as a gRPC API, but where possible we will continue to ensure that teleport does not make unnecessary
network round trips.

We will also establish a set of conventions intended to make the transition to using the new system as easy as possible,
and to ensure that it is as difficult as possible to accidentally misuse it either by misunderstanding or oversight.

### API

The PDP/PEP boundary will be represented as a gRPC API, with methods for the different kinds of decisions. Because teleport
as a whole is an integrated PDP/PEP, our decisions are often much more nuanced than a simple allow/deny. This is one of
teleport's greatest strengths, as this allows us to provide very granular controls over actions. It also means that many
of the decisions we make are very verbose. Take ssh access for example: an 'allow' decision for ssh access must include
parameters for agent forwarding, port forwarding, concurrent connection limits, BPF events, expired cert disconnect, locking,
file copying, etc. For this reason we will eschew any attempt at creating a "unified" method for decisions in favor of
implementing custom methods and types intended to meet the specific needs of specific decisions.

In order to ensure that the PDP API remains simple and usable we will establish a set of conventions for how decision methods
and types should be handled:

- All "allow" decisions will be structured as conditional allows, pending the application of limits/parameters. All allow decisions
will be communicated by a type of the form `<Action>Permit`, which will encode all limits/parameters to be applied by the PEP. Much of
the agent-side teleport logic that previously would have worked with some combination of certificate identity and role set will instead
take the appropriate permit structure as input. Permits will *always* be passed by pointer, not value, as a precaution against zero
value bugs resulting in unintended access.

- Methods will conventionally take the form `rpc Evaluate<Action>(<Action>Request) returns (<Action>Permit)`. For example,
the method for evaluating server access would read `rpc EvaluateSSHAccess(SSHAccessRequest) returns (SSHAccessPermit)`.
Note that this deviates somewhat from common PDP conventions in that we don't return a decision object wrapping the permit. Any result
other than allow must be communicated as an error.

- Common `RequestMetadata` and `PermitMetadata` types will be provided and should be included in the request and permit
respectively. These types will contain common fields like teleport version of PDP/PEP, dry run flag, etc.

- A standard `Resource` reference types with fields like  `kind` and `name` will be provided and should be prefered in
those APIs where specifying individual resources or sets of resources is necessary. PDPs will use this reference to load
the appropriate resource from local caches.

- Two standard `Identity` types will be provided for use in communicating caller identity within decision requests. One
corresponding to standard teleport TLS cert format, and one to the standard SSH cert format. Eventually, we would like
the PDP to be able to make decisions based on a standardized user/subject reference similar to the `Resource`
reference, with most fields currently stored in the certificate instead determined at runtime. In practice, moving away from
teleport's dependence on complex certificate-based identities is non-trivial. In order to support an iterative transition, we will
reimplement our existing identity types as a protobuf type and use that as the standard subject identifier. Individual fields will be
transitioned from being cert-bound to being determined PDP-side in an iterative manner.

Here is a truncated example of the PDP API for server access:

```protobuf
service DecisionService {

  rpc EvaluateSSHAccess(SSHAccessRequest) returns (SSHAccessPermit);

  // ...
}

message SSHAccessRequest {
  RequestMetadata metadata = 1;

  SSHIdentity user = 2;

  Resource server = 3;

  string login = 4;
}

message SSHAccessPermit {
  PermitMetadata metadata = 1;

  repeated string logins = 2;

  bool forward_agent = 3;

  google.protobuf.Duration max_session_ttl = 4;

  bool port_forwarding = 5;

  int64 client_idle_timeout = 6;

  bool disconnect_expired_cert = 7;

  repeated string bpf = 8;

  bool x11_forwarding = 9;

  int64 max_connections = 10;

  // ... (there's a lot more that needs to go here)
}

message SSHIdentity {
  string teleport_user = 1;
  
  string impersonator = 2;
  
  string login = 3;

  repeated string unmapped_roles = 4;

  google.protobuf.Timestamp cert_valid_before = 5;

  string route_to_cluster = 6;

  // ... (there's a lot more that needs to go here)
}
```

With the above API, most server access enforcement logic can be reworked to respect the PDP/PEP boundary with minimal changes.
`AccessChecker.EnhancedRecordingSet()` becomes `permit.Bpf`, `AccessChecker.GetHostSudoerts(server)` becomes
`permit.HostSudoers`, etc.

### Phased Implementation

The switch to PDP will be accomplished in two distinct phases. An "refactor" phase, and a "relocate" phase.

#### Refactor Phase

The refactor phase will consist of initial implementation of the PDP API, without any outward changes that might affect
cross-version compatibility. In order to seamlessly transition from decisions being made at agents to decisions being
made at the control plane, the implementation of the PDP logic will be polymorphic over two possible operating modes
depending on where it is running.

On the control plane, the PDP will have access to the entire set of teleport users and roles. This variant will be able
to serve as the backing of the gRPC API, and make internal decisions as if it were a remote PDP. This PDP will be a shared
service initialized once at process startup.

On agents, a local ephemeral PDP will be able to be initialized with a fake user store only knowing the identity of the
user making the request. This variant of the PDP will need to be initialized for each incoming request, and will be able to
only serve requests related to the user for which it was initialized (much like the `AccessChecker` today).

Common PDP logic will be shared between the two variants, meaning that so long as we need to support the agent-local PDP,
the control-plane PDP won't be able to make use of any state that cannot be derived from a combination of local agent cache
and user certificate identity. The APIs of both local and remote PDP will conform to the same interface so that any logic
relying upon the PDP can abstract over local and remote implementations.

During this phase we want to start moving as many elements as possible from being part of the certificate-derived identity to
being calculated on-demand within the PDP itself. As a prerequisite to this, we will need to implement a custom access request
cache, as access requests have a signficant impact on many identity fields. A custom cache will be necessary in order to handle
the fact that access requests are often used immediately upon approval, before approval can be relied upon to be propagated
to caches naturally.

The intent of the above work is to allow us to make the vast majority of necessary code changes in a totally backwards-compatible
manner and to backport them to all active release branches. This will ensure that ongoing work related to policy enforcement
doesn't need to reimplement the same logic for the old model when backporting. Though its worth considering trying to land
all major changes prior to a testplan so that we can get robust manual testing of all major features with the changes in
place before backporting.

To ensure that we've fully broken agent dependence on roles outside of the PDP, we will split the agent's local cache up
s.t. all components other than the PDP will be statically prevented from accessing roles.

During this phase we will also implement tctl commands for directly invoking the DecisionService gRCP API on the control
plane. This will provide a powerful debugging and auditing tool for superusers and developers who want deeper insight into how
teleport makes decisions.

#### Relocate Phase

The relocate phase will see us actually transition to all access-control decisions being made at the control plane. The method
of relocation will depend on the kind of teleport agent. For tunnel agents, we will start performing denial at the control plane
itself, and allowed dials will have the Permit object forwarded to the agent as part of the dial. For direct dial agents, the
agent will call-back into the control plane to get a decision.

We intend to open a second follow-up RFD when we are closer to the relocate phase in order to explore it in more detail, but the
highlights are these:

- Agent reverse tunnel protocols will need to be updated to allow a trusted and replay-resistant permit message to be sent from proxy to
agent as part of an incoming dial.

- The mechanism by which trusted cluster dials are performed will need to be reworked s.t. all routing decisions are made at
the leaf proxy rather than the root proxy since access-control decisions will now get made as a part of routing. This will include
sending sideband information about user identity when forwarding the dial from root proxy to leaf proxy in order to ensure that
the leaf proxy can correctly map the identity without relying on user certificate information.

- One major version of buffer will be needed during which agents still cache roles locally and roles/traits are still applied to certs
so that we can fallback to local decisions when talking to outdated control plane elements.

- Agentless/openssh nodes are already using a "decisions at control plane" model, but they still depend on logins being present in
certificates. We may want to consider a mechanism for eliminating logins at the same time (e.g. by lazily provisioning certs containing
only the target login at the proxy).

### Questions/Alternatives/Ongoing Research

- We've gone back and forth a bit on whether to return bare permits s.t. the resulting client method returns `(*Permit, error)` or
to use an enclosing response/decision message with a oneof for allow/deny cases. The former is more inline with how existing access
control decisions are made and results in slightly cleaner code, but the latter may be useful if we ever decide to start doing more
complex denial cases (e.g. for MFA flows or verbose denials for use by visualization tools). We're currently leaning toward bare
permits since all the good arguments for enclosed responses are speculative at this point. One of the design goals of this system
is to let us iterate faster, and that includes adding new methods with new conventions as the need arises.

- Some protocols (e.g. k8s) don't map as cleanly to a single decision being made in advance at dial time. We are still looking into our
options there and plans related to such protocols are subject to change. SSH being the most widely used protocol, and the protocol that
incurs by far the most role replication load, has been the focus of our investigation thus far. It may end up being that
some protocols' permits end up being something more akin to a compiled "resource selector" that can be sent as part of a dial. Breaking
the agent's dependence on role replication while still leaving a lot of the actual per-resource decision making agent-side.
