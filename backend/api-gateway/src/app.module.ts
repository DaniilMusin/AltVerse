import { Module }        from '@nestjs/common';
import { GraphQLModule } from '@nestjs/graphql';
import { PrismaModule }  from './prisma/prisma.module';
import { AuthModule }    from './modules/auth/auth.module';
import { PlayerModule }  from './modules/player/player.module';

@Module({
  imports: [
    PrismaModule,
    AuthModule,
    PlayerModule,
    GraphQLModule.forRoot({
      autoSchemaFile: 'schema.gql',
      context: ({ req }) => ({ headers: req.headers }),
    }),
  ],
})
export class AppModule {}
