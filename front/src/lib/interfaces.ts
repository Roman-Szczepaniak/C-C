// User interface
export interface User {
    id: string;
    email: string;
    firstname: string;
    lastname: string;
    createdAt: Date;
    updatedAt: Date;
}

// Player interface
export interface Player {
    id: number;
    user_id: number;
    party_id: number;
    name: string;
    level: number;
    initiative?: number;
    hp: number;
    hp_max: number;
    xp?: boolean;
}

// Party interface
export interface Party {
    id: number;
    created_at?: Date;
    updated_at?: Date;
}

// Monster interface
export interface Monster {
    id: number;
    name: string;
    alignment: string;
    size: string;
    type: string;
    environment: string;
    status: string;
    cr: string;
    card_id?: number;
}

// Encounter interface
export interface Encounter {
    id: number;
    party_id: number;
    difficulty: string;
    type: string;
    total_xp?: string;
}

// Card interface
export interface Card {
    id: number;
    description: string;
    ca: number;
    pv: number;
    speed: string;
    strength: string;
    dexterity: string;
    constitution: string;
    intelligence: string;
    wisdom: string;
    charisma: string;
    save_throws?: string;
    damage_resist?: string;
    damage_immune?: string;
    condition_immune?: string;
    senses?: string;
    languages?: string;
}

// Action interface
export interface Action {
    id: number;
    name: string;
    description: string;
}
