import {z} from 'zod';
import {UserSchema} from './user';

export const StaffUsers = z.array(UserSchema);
export type StaffUsers = z.infer<typeof StaffUsers>;
