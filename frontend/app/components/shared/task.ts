module tasks.shared{

    export interface Task{
        id: number;
        text: string;
        progress: [{period: Period, progress: number}]
        recurring: boolean;
        totalWork: number;
        firstPeriod: Period;
        lastPeriod?: Period;
        type: TaskType;
        deleted?: boolean;
    }

    export enum TaskType {
        MUST,
        WANT_TO,
    }

    export class Period {
        week: number;
        year: number;
        s: string;

        constructor(s: string);
        constructor(w: number, y: number);
        constructor(a: any, y?: number){
            if ( typeof a === 'string' ){
                var spl = a.split("-");
                this.week = parseInt(spl[0]);
                this.year = parseInt(spl[1]);
                this.s = a;
            } else {
                this.week = a;
                this.year = y;
                this.s = this.week.toString() + "-" + this.year;
            }
        }

        toString(): string {
            return this.s;
        }
    }
}