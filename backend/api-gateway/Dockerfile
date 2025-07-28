# -------- build --------
FROM node:20-bookworm AS deps
WORKDIR /app
COPY backend/api-gateway/package.json backend/api-gateway/pnpm-lock.yaml ./
RUN corepack enable && pnpm install --frozen-lockfile

FROM deps AS build
COPY backend/api-gateway .
RUN pnpm run build

# -------- run ---------
FROM gcr.io/distroless/nodejs20-debian12:nonroot
WORKDIR /app
COPY --from=build /app/dist ./dist
COPY --from=deps /app/node_modules ./node_modules
ENV NODE_ENV=production
CMD ["node", "dist/main.js"]
