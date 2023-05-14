import {z} from 'zod';

export const LoginDeviceScheme = z.object({
  id: z.number(),

  device: z.string().nullable(),
  os: z.string().nullable(),
  browser: z.string().nullable(),
  is_mobile: z.boolean().nullable(),
  ip: z.string(),

  is_current: z.boolean(),

  created: z.string().datetime({offset: true}),
});
export type LoginDevice = z.infer<typeof LoginDeviceScheme>;

export const LoginDeviceListScheme = z.array(LoginDeviceScheme);
export type LoginDeviceList = z.infer<typeof LoginDeviceListScheme>;
