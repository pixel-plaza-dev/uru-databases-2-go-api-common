import { PickType } from '@nestjs/swagger';
import { UserDTO } from './user.dto';

export class UserSendEmailVerificationTokenDTO extends PickType(UserDTO, [
  'email',
] as const) {}
