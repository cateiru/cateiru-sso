import {z} from 'zod';

export const PublicAuthenticationLoginSessionSchema = z.object({
  login_session_token: z.string(),
  limit_date: z.string().datetime({offset: true}),
});
export type PublicAuthenticationLoginSession = z.infer<
  typeof PublicAuthenticationLoginSessionSchema
>;

export const PublicAuthenticationRequestSchema = z.object({
  client_id: z.string(),
  client_name: z.string(),
  client_description: z.string().nullable(),
  image: z.string().url().nullable(),

  org_name: z.string().nullable(),
  org_image: z.string().url().nullable(),
  org_member_only: z.boolean(),

  scopes: z.array(z.string()).nullable(),
  redirect_uri: z.string().url(),
  response_type: z.string(),

  register_user_name: z.string(),
  register_user_image: z.string().nullable(),

  prompts: z.array(z.string()),

  login_session: PublicAuthenticationLoginSessionSchema.optional(),
});
export type PublicAuthenticationRequest = z.infer<
  typeof PublicAuthenticationRequestSchema
>;
