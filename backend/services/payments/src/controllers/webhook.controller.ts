import Stripe from 'stripe';

export async function handleStripeWebhook(
  stripe: Stripe,
  signature: string,
  rawBody: Buffer,
) {
  const event = stripe.webhooks.constructEvent(
    rawBody,
    signature,
    process.env.STRIPE_WEBHOOK_SECRET!,
  );

  switch (event.type) {
    case 'checkout.session.completed': {
      const session = event.data.object as Stripe.Checkout.Session;
      // TODO: пометить invoice paid, начислить бустер-NFT
      break;
    }
    default:
      console.log(`Unhandled event ${event.type}`);
  }
}
