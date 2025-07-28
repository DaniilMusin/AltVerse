import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { CreatePlayerInput } from './dto/create-player.input';

@Injectable()
export class PlayerService {
  constructor(private readonly prisma: PrismaService) {}

  async create(data: CreatePlayerInput) {
    return this.prisma.player.create({ data });
  }

  async findById(id: string) {
    return this.prisma.player.findUnique({ where: { id } });
  }

  async leaderboard(top = 100) {
    return this.prisma.player.findMany({
      orderBy: { stats: { path: ["score"], sort: "desc" } },
      take: top,
      select: { id: true, username: true, stats: true },
    });
  }
}
