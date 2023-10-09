import z from 'zod';

export const publicAuthenticationRequestSchema = z.object({
  client_id: z.string(),
  client_name: z.string(),
  client_description: z.string().nullable(),
  image: z.string().url().nullable(),

  org_name: z.string().nullable(),
  org_image: z.string().url().nullable(),

  scopes: z.array(z.string()),
  redirect_uri: z.string().url(),
  response_type: z.string(),
});
export type PublicAuthenticationRequest = z.infer<
  typeof publicAuthenticationRequestSchema
>;
