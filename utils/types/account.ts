import {z} from 'zod';

export const AccountUserSchema = z.object({
  user_name: z.string(),
  id: z.string(),
  avatar: z.string().optional(),
});
export type AccountUser = z.infer<typeof AccountUserSchema>;

export const AccountUserListSchema = z.array(AccountUserSchema);
export type AccountUserList = z.infer<typeof AccountUserListSchema>;

export const AccountCertificatesSchema = z.object({
  password: z.boolean(),
  otp: z.boolean(),
  otp_modified: z.string().datetime({offset: true}).nullable(),
});
export type AccountCertificates = z.infer<typeof AccountCertificatesSchema>;

export const AccountOTPPublicSchema = z.object({
  otp_session: z.string(),
  public_key: z.string(),
});
export type AccountOTPPublic = z.infer<typeof AccountOTPPublicSchema>;

export const AccountWebAuthnDeviceSchema = z.object({
  id: z.number(),

  device: z.string().nullable(),
  os: z.string().nullable(),
  browser: z.string().nullable(),
  is_mobile: z.boolean().nullable(),
  ip: z.string().nullable(),

  created: z.string().datetime({offset: true}),
});
export type AccountWebAuthnDevice = z.infer<typeof AccountWebAuthnDeviceSchema>;

export const AccountWebAuthnDevicesSchema = z.array(
  AccountWebAuthnDeviceSchema
);
export type AccountWebAuthnDevices = z.infer<
  typeof AccountWebAuthnDevicesSchema
>;
