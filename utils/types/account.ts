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
