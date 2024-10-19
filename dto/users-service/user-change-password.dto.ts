import {IsNotEmpty, IsString, MaxLength, MinLength} from 'class-validator';
import {ApiProperty, PickType} from '@nestjs/swagger';
import {UserDTO} from './user.dto';

export class UserChangePasswordDTO extends PickType(UserDTO, [
    'password',
    'confirmPassword',
] as const) {
    @IsString()
    @IsNotEmpty()
    @MinLength(12)
    @MaxLength(64)
    @ApiProperty()
    readonly currentPassword: string;
}
