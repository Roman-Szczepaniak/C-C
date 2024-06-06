import type { PageLoad } from './$types';
import { get } from '$lib/api';

export const load = (async ({ params, depends }) => {
    depends('app:monster');

    const monsterResponse = await get(`/monsters/${params.monster_id}`);
    let cardResponse = null;

    if (monsterResponse.card_id) {
        cardResponse = await get(`/cards/${monsterResponse.card_id}`);
        if (cardResponse.actions) {
            cardResponse.actions = await Promise.all(cardResponse.actions.map(async (action) => {
                const actionData = await get(`/actions/${action.id}`);
                return { ...action, ...actionData };
            }));
        }
    }

    return {
        title: 'Monster',
        monster: {
            ...monsterResponse,
            card: cardResponse
        }
    };
}) satisfies PageLoad;