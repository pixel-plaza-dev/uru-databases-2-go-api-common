import {IntersectionType, PickType} from '@nestjs/swagger';
import {UserDTO} from './user.dto';
import {UserVerifyEmailDTO} from "./user-verify-email.dto";

export class UserResetPasswordDTO extends IntersectionType(
    PickType(UserDTO, [
        'password',
        'confirmPassword',
    ] as const),
    PickType(UserVerifyEmailDTO, ['uuid'] as const)) {
}
