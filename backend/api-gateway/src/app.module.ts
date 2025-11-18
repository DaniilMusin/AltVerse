import { Module }        from '@nestjs/common';
import { GraphQLModule } from '@nestjs/graphql';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { PrismaModule }  from './prisma/prisma.module';
import { AuthModule }    from './modules/auth/auth.module';
import { PlayerModule }  from './modules/player/player.module';

@Module({
  imports: [
    PrismaModule,
    AuthModule,
    PlayerModule,
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      autoSchemaFile: 'schema.gql',
      context: ({ req }) => ({ headers: req.headers }),
    }),
  ],
})
export class AppModule {}
