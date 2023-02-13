/**
 * Copyright 2020-2022 Gravitational, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import React from 'react';
import { render, fireEvent, screen, waitFor } from 'design/utils/testing';
import { privateKeyEnablingPolicies } from 'shared/services/consts';

import auth from 'teleport/services/auth/auth';
import cfg from 'teleport/config';

import Login from './Login';

const assignMock = jest.fn();

delete window.location;
(window.location as any) = { assign: assignMock };

afterEach(() => {
  assignMock.mockClear();
});

beforeEach(() => {
  jest.restoreAllMocks();
});

test('basic rendering', () => {
  render(<Login />);

  // test rendering of logo and title
  expect(screen.getByRole('img')).toBeInTheDocument();
  expect(screen.getByText(/sign in to teleport/i)).toBeInTheDocument();
});

test('login with redirect', async () => {
  jest.spyOn(auth, 'login').mockResolvedValue(null);

  render(<Login />);

  // fill form
  const username = screen.getByPlaceholderText(/username/i);
  const password = screen.getByPlaceholderText(/password/i);
  fireEvent.change(username, { target: { value: 'username' } });
  fireEvent.change(password, { target: { value: '123' } });

  // test login and redirect
  fireEvent.click(screen.getByText('Sign In'));
  await waitFor(() => {
    expect(auth.login).toHaveBeenCalledWith('username', '123', '');
  });

  expect(assignMock).toHaveBeenCalledWith('http://localhost/web');
});

test('login with SSO', async () => {
  jest.spyOn(cfg, 'getAuth2faType').mockImplementation(() => 'otp');
  jest.spyOn(cfg, 'getPrimaryAuthType').mockImplementation(() => 'sso');
  jest.spyOn(cfg, 'getAuthProviders').mockImplementation(() => [
    {
      displayName: 'With Github',
      type: 'github',
      name: 'github',
      url: '/github/login/web',
    },
  ]);

  render(<Login />);

  console.log(screen.getByText('With Github'));

  // test login pathways
  fireEvent.click(screen.getByText('With Github'));

  expect(assignMock).toHaveBeenCalledWith(
    'http://localhost/github/login/web?connector_id=github&redirect_url=http%3A%2F%2Flocalhost%2Fweb'
  );
});

test('login with private key policy enabled through cluster wide', () => {
  jest
    .spyOn(cfg, 'getPrivateKeyPolicy')
    .mockImplementation(() => 'hardware_key');

  render(<Login />);

  expect(screen.queryByPlaceholderText(/username/i)).not.toBeInTheDocument();
  expect(screen.getByText(/login disabled/i)).toBeInTheDocument();
});

test('login with private key policy enabled through role setting', async () => {
  // Just needs any of these enabling keywords in error message
  jest
    .spyOn(auth, 'login')
    .mockRejectedValue(new Error(privateKeyEnablingPolicies[0]));

  render(<Login />);

  // Fill form.
  const username = screen.getByPlaceholderText(/username/i);
  const password = screen.getByPlaceholderText(/password/i);
  fireEvent.change(username, { target: { value: 'username' } });
  fireEvent.change(password, { target: { value: '123' } });

  // Test logging in with private key error return renders private policy error.
  fireEvent.click(screen.getByText('Sign In'));
  await waitFor(() => {
    expect(auth.login).toHaveBeenCalledWith('username', '123', '');
  });

  expect(screen.queryByPlaceholderText(/username/i)).not.toBeInTheDocument();
  expect(screen.getByText(/login disabled/i)).toBeInTheDocument();
});
