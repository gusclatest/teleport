/**
 * Teleport
 * Copyright (C) 2024 Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import React from 'react';

import Dialog from 'design/Dialog';

import { http, HttpResponse, delay } from 'msw';

import { createTeleportContext } from 'teleport/mocks/contexts';
import { ContextProvider } from 'teleport/index';

import cfg from 'teleport/config';

import {
  DeviceType,
  DeviceUsage,
  getMfaRegisterOptions,
} from 'teleport/services/mfa';

import {
  AddAuthDeviceWizardStepProps,
  CreateDeviceStep,
  SaveDeviceStep,
} from './AddAuthDeviceWizard';
import { ReauthenticateStep } from './ReauthenticateStep';

export default {
  title: 'teleport/Account/Manage Devices/Add Device Wizard',
  decorators: [
    Story => {
      const ctx = createTeleportContext();
      return (
        <ContextProvider ctx={ctx}>
          <Dialog open={true} dialogCss={() => ({ width: '650px' })}>
            <Story />
          </Dialog>
        </ContextProvider>
      );
    },
  ],
};

export function Reauthenticate() {
  return <ReauthenticateStep {...stepProps} />;
}

export function ReauthenticateLimitedOptions() {
  return (
    <ReauthenticateStep
      {...stepProps}
      mfaChallengeOptions={[{ value: 'totp', label: 'Authenticator App' }]}
    />
  );
}

export function ReauthenticateNoOptions() {
  return <ReauthenticateStep {...stepProps} />;
}

export function CreatePasskey() {
  return <CreateDeviceStep {...stepProps} usage="passwordless" />;
}

export function CreateMfaHardwareDevice() {
  return (
    <CreateDeviceStep {...stepProps} usage="mfa" newMfaDeviceType="webauthn" />
  );
}

export function CreateMfaAppQrCodeLoading() {
  return (
    <CreateDeviceStep {...stepProps} usage="mfa" newMfaDeviceType="totp" />
  );
}
CreateMfaAppQrCodeLoading.parameters = {
  msw: {
    handlers: [
      http.post(
        cfg.getMfaCreateRegistrationChallengeUrl('privilege-token'),
        async () => await delay('infinite')
      ),
    ],
  },
};

export function CreateMfaAppQrCodeFailed() {
  return (
    <CreateDeviceStep {...stepProps} usage="mfa" newMfaDeviceType="totp" />
  );
}
CreateMfaAppQrCodeFailed.parameters = {
  msw: {
    handlers: [
      http.post(
        cfg.getMfaCreateRegistrationChallengeUrl('privilege-token'),
        () =>
          HttpResponse.json(
            {
              error: { message: 'Whoops, something went wrong.' },
            },
            { status: 500 }
          )
      ),
    ],
  },
};

const dummyQrCode =
  'iVBORw0KGgoAAAANSUhEUgAAAB0AAAAdAQMAAABsXfVMAAAABlBMVEUAAAD///+l2Z/dAAAAAnRSTlP//8i138cAAAAJcEhZcwAACxIAAAsSAdLdfvwAAABrSURBVAiZY/gPBAxoxAcxh3qG71fv1zN8iQ8EEReBRACQ+H4ZKPZBFCj7/3v9f4aPU9vqGX4kFtUzfG5mBLK2aNUz/PM3AsmqAk2RNQTquLYLqDdG/z/QlGAgES4CFLu4GygrXF2Pbi+IAADZqFQFAjXZWgAAAABJRU5ErkJggg==';

export function CreateMfaApp() {
  return (
    <CreateDeviceStep {...stepProps} usage="mfa" newMfaDeviceType="totp" />
  );
}
CreateMfaApp.parameters = {
  msw: {
    handlers: [
      http.post(
        cfg.getMfaCreateRegistrationChallengeUrl('privilege-token'),
        () => HttpResponse.json({ totp: { qrCode: dummyQrCode } })
      ),
    ],
  },
};

export function SavePasskey() {
  return <SaveDeviceStep {...stepProps} usage="passwordless" />;
}

export function SaveMfaHardwareDevice() {
  return (
    <SaveDeviceStep {...stepProps} usage="mfa" newMfaDeviceType="webauthn" />
  );
}

export function SaveMfaAuthenticatorApp() {
  return <SaveDeviceStep {...stepProps} usage="mfa" newMfaDeviceType="totp" />;
}

const stepProps: AddAuthDeviceWizardStepProps = {
  // StepComponentProps
  next: () => {},
  prev: () => {},
  hasTransitionEnded: true,
  stepIndex: 0,
  flowLength: 1,
  refCallback: () => {},

  // Shared props
  privilegeToken: 'privilege-token',
  newMfaDeviceType: 'webauthn',
  onClose: () => {},
  onSuccess: () => {},
  usage: 'passwordless' as DeviceUsage,

  // Reauth props
  reauthAttempt: null,
  clearReauthAttempt: () => {},
  mfaChallengeOptions: getMfaRegisterOptions('on'),
  submitWithMfa: async (mfaType?: DeviceType, otpCode?: string) => {},

  // Create props
  mfaRegisterOptions: null,
  onDeviceCreated: (c: Credential) => {},
  onNewMfaDeviceTypeChange: (d: DeviceType) => {},

  // Save props
  credential: { id: 'cred-id', type: 'public-key' },
};
