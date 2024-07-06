import {z} from 'zod';

export const OidcParamsSchema = z.object({
  scope: z.string(),
  response_type: z.string(),
  client_id: z.string(),
  redirect_uri: z.string(),
  state: z.string(),

  response_mode: z.string().nullable(),
  nonce: z.string().nullable(),
  display: z.string().nullable(),
  prompt: z.string().nullable(),
  max_age: z.string().nullable(),
  ui_locales: z.string().nullable(),
  id_token_hint: z.string().nullable(),
  login_hint: z.string().nullable(),
  acr_values: z.string().nullable(),
});
export type OidcParams = z.infer<typeof OidcParamsSchema>;
