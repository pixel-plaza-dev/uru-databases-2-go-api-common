import {PickType} from '@nestjs/swagger';
import {UserDTO} from './user.dto';

export class UserForgotPasswordDTO extends PickType(UserDTO, [
    'username',
    'email',
] as const) {
}
