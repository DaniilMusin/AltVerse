FROM node:20-bookworm AS deps
WORKDIR /app
COPY backend/services/payments/package*.json ./
RUN corepack enable && pnpm install --frozen-lockfile

FROM deps AS build
COPY backend/services/payments .
RUN pnpm run build

FROM gcr.io/distroless/nodejs20-debian12:nonroot
WORKDIR /app
COPY --from=build /app/dist ./dist
COPY --from=deps /app/node_modules ./node_modules
CMD ["node", "dist/main.js"]
