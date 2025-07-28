import Fastify from 'fastify';
import Stripe from 'stripe';
import { handleStripeWebhook } from './controllers/webhook.controller';

const app    = Fastify();
const stripe = new Stripe(process.env.STRIPE_KEY!, { apiVersion: '2023-10-16' });

app.post('/webhook/stripe', async (request, reply) => {
  const sig = request.headers['stripe-signature'] as string;
  const raw = (request as any).rawBody;
  await handleStripeWebhook(stripe, sig, raw);
  reply.status(200).send('ok');
});

app.listen({ port: 8080, host: '0.0.0.0' });
