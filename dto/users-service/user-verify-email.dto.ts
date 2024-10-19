import {IsAlphanumeric, IsNotEmpty} from "class-validator";
import {ApiProperty} from "@nestjs/swagger";

export class UserVerifyEmailDTO {
    @IsAlphanumeric()
    @IsNotEmpty()
    @ApiProperty()
    readonly uuid: string;
}