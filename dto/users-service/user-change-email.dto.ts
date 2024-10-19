import {IsEmail, IsNotEmpty} from 'class-validator';
import {ApiProperty, PickType} from '@nestjs/swagger';
import {UserDTO} from './user.dto';

export class UserChangeEmailDTO extends PickType(UserDTO, [
    'email',
] as const) {
    @IsEmail()
    @IsNotEmpty()
    @ApiProperty()
    readonly newEmail: string;
}
