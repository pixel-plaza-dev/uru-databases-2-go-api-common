import {PickType} from '@nestjs/swagger';
import {UserDTO} from '../users-service/user.dto';

export class UserAuthLoginDTO extends PickType(UserDTO, [
    'username',
    'password',
] as const) {
}
