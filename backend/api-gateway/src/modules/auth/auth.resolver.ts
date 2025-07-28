import { Args, Mutation, Resolver } from '@nestjs/graphql';
import { AuthService } from './auth.service';

@Resolver()
export class AuthResolver {
  constructor(private readonly auth: AuthService) {}

  @Mutation(() => String)
  async login(
    @Args('username') username: string,
    @Args('password') password: string,
  ) {
    const user = await this.auth.validateUser(username, password);
    if (!user) throw new Error('Invalid credentials');
    const { accessToken } = await this.auth.login(user);
    return accessToken;
  }
}
