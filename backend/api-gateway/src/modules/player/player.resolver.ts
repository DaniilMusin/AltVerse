import { Args, Int, Mutation, Query, Resolver } from '@nestjs/graphql';
import { PlayerService } from './player.service';
import { Player } from './player.model';
import { CreatePlayerInput } from './dto/create-player.input';

@Resolver(() => Player)
export class PlayerResolver {
  constructor(private readonly svc: PlayerService) {}

  @Mutation(() => Player)
  createPlayer(@Args('data') data: CreatePlayerInput) {
    return this.svc.create(data);
  }

  @Query(() => Player, { nullable: true })
  player(@Args('id') id: string) {
    return this.svc.findById(id);
  }

  @Query(() => [Player])
  leaderboard(@Args('take', { type: () => Int, defaultValue: 100 }) take: number) {
    return this.svc.leaderboard(take);
  }
}
