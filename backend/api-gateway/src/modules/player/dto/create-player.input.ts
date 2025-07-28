import { Field, InputType } from '@nestjs/graphql';
import { IsAlphanumeric, Length } from 'class-validator';

@InputType()
export class CreatePlayerInput {
  @Field()
  @IsAlphanumeric()
  @Length(3, 16)
  username: string;

  @Field({ nullable: true })
  wallet?: string;
}
