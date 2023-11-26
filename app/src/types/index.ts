import type { NumberLiteralType } from "typescript";

interface User {
    id: number;
    username: string;
    email: string;
    password: string;
    created_at: string;
    updated_at: string;
}

interface Project {
    id: number;
    name: string;
    description: string;
    created_at: string;
    updated_at: string;
    user_id: number;
    lists: List[];
}

interface List {
    id: number;
    name: string;
    position: number;
    createdAt: string;
    updatedAt: string;
    project_id: number;
    tasks: Task[];
}

interface Task {
    id: number;
    name: string;
    description: string;
    created_at: string;
    updated_at: string;
    list_id: number;
}

interface Alert {
    type: string;
    message: string;
}

type AlertType = 'success' | 'error' | 'warning' | 'info';

export type { User, Alert, AlertType, Project, List, Task };
