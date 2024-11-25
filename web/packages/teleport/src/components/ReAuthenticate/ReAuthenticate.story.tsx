/**
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
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

import { Props, ReAuthenticate } from './ReAuthenticate';
import { DeviceType, getMfaRegisterOptions } from 'teleport/services/mfa';

export default {
  title: 'Teleport/ReAuthenticate',
};

export const Loaded = () => <ReAuthenticate {...props} />;

export const Processing = () => (
  <ReAuthenticate {...props} attempt={{ status: 'processing' }} />
);

export const Failed = () => (
  <ReAuthenticate
    {...props}
    attempt={{ status: 'failed', statusText: 'an error has occurred' }}
  />
);

const props: Props = {
  attempt: { status: '' },
  clearAttempt: () => null,
  getMfaChallenge: () => null,
  getMfaChallengeOptions: async () => getMfaRegisterOptions('on'),
  submitWithMfa: (mfaType?: DeviceType, totp_code?: string) => null,
  onClose: () => null,
};
