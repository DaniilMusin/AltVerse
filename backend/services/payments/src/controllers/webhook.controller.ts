import Stripe from 'stripe';

export async function handleStripeWebhook(stripe: Stripe, sig: string, raw: Buffer) {
  // TODO: implement webhook verification
  return;
}
