import { Injectable } from '@nestjs/common';
import { PassportStrategy } from '@nestjs/passport';
import { Strategy } from 'passport-jwt';

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
  constructor() {
    super({
      jwtFromRequest: (req) => req.headers.authorization ?? '',
      secretOrKey: process.env.JWT_SECRET,
    });
  }

  validate() { return {}; }
}
