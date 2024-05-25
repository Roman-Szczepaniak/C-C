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
    name: string;
    level: number;
    initiative?: number;
    hp: number;
    hp_max: number;
    xp?: boolean;
    accountId: number;
    partyId: number;
}

// Party interface
export interface Party {
    id: number;
    createdAt?: Date;
    updatedAt?: Date;
}

// Monster interface
export interface Monster {
    id: number;
    name: string;
    alignment?: string;
    size: string;
    type: string;
    cr: string;
    monsterId: number;
}

// Encounter interface
export interface Encounter {
    id: number;
    difficulty: string;
    type: string;
    totalXp?: string;
    partyId: number;
}

// Card interface
export interface Card {
    id: number;
    description?: string;
    ca?: number;
    pv?: number;
    speed?: string;
    action?: string;
    strength: string;
    dexterity: string;
    constitution: string;
    intelligence: string;
    wisdom: string;
    charisma: string;
}
