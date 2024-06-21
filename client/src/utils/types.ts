export interface User {
    id: number;
    name: string;
    avatarUrl: string;
}

export interface HrNote {
    timestamp: number;
    body: string;
}

export interface HrNotes {
    userId: number;
    notes: HrNote[];
}