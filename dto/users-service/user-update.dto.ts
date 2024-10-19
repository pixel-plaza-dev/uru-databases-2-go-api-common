import {PartialType, PickType} from '@nestjs/swagger';
import {UserDTO} from './user.dto';

export class UserUpdateDTO extends PartialType(
    PickType(UserDTO, [
        'address',
        'phone',
        'birthDate',
        'firstName',
        'lastName',
    ] as const),
) {
}
