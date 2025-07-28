import Stripe from 'stripe';

export class StripeAdapter {
  constructor(private readonly client: Stripe) {}

  async createCheckout(amountCents: number, playerId: string) {
    return this.client.checkout.sessions.create({
      mode: 'payment',
      metadata: { playerId },
      line_items: [
        { price_data: { currency: 'usd', unit_amount: amountCents, product_data: { name: 'Premium Pass' } }, quantity: 1 },
      ],
      success_url: process.env.SUCCESS_URL!,
      cancel_url: process.env.CANCEL_URL!,
    });
  }
}
