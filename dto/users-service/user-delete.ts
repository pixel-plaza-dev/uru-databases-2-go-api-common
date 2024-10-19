import {PickType} from '@nestjs/swagger';
import {UserDTO} from './user.dto';

export class UserDeleteDTO extends PickType(UserDTO, [
    'username',
    'password',
    'confirmPassword',
] as const) {
}
