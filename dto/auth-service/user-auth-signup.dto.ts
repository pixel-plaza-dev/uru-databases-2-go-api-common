import { PickType } from '@nestjs/swagger';
import { UserDTO } from '../users-service/user.dto';

export class UserAuthSignupDTO extends PickType(UserDTO, [
  'email',
  'username',
  'password',
  'confirmPassword',
  'firstName',
  'lastName',
] as const) {}
