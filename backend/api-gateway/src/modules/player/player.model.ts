import { Field, ObjectType } from '@nestjs/graphql';

@ObjectType()
export class Player {
  @Field()
  id: string;

  @Field()
  username: string;

  @Field({ nullable: true })
  wallet?: string;

  @Field(() => String, { nullable: true })
  stats?: any;
}
