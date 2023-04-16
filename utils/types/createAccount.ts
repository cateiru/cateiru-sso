import z from 'zod';

export const CreateAccountRegisterEmailResponseSchema = z.object({
  register_token: z.string(),
});
export type CreateAccountRegisterEmailResponse = z.infer<
  typeof CreateAccountRegisterEmailResponseSchema
>;
