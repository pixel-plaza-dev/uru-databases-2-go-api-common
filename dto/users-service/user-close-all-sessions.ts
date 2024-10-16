import { PickType } from '@nestjs/swagger';
import { UserDTO } from './user.dto';

export class UserCloseAllSessionsDTO extends PickType(UserDTO, [
  'password',
] as const) {}
