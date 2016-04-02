module tasks.shared{

    export interface Task{
        id: number;
        userID: number;
        text: string;
        workload: number;
        workDone: number;
        recurring: boolean;
	    unit: string;
	    mustDo: boolean;
	    periodStart: Date;
    }
}