interface LedgerEntry {
    journalId: string;
    entryId: string;
    amount: number;
    type: 'DEBIT' | 'CREDIT';
    asset: string;
    productLine: 'EQUITY' | 'CRYPTO' | 'DERIVATIVE'; 
    accountId: string;
    timestamp: number;
}

function generateUniqueId(): string {
    return Date.now().toString(36) + Math.random().toString(36).substring(2);
}

class Ledger {
    entries: LedgerEntry[] = []; 

    constructor() {}

    recordTransaction(
        debitAccountId: string, 
        creditAccountId: string, 
        amount: number, 
        asset: string, 
        productLine: LedgerEntry['productLine']
    ): { creditEntry: LedgerEntry, debitEntry: LedgerEntry } {
        const timestamp = Date.now();
        const journalId = generateUniqueId();
        
        const creditEntry: LedgerEntry = {
            entryId: generateUniqueId(),
            journalId,
            type: 'CREDIT',
            amount,
            accountId: creditAccountId,
            asset,
            productLine,
            timestamp
        };

        const debitEntry: LedgerEntry = {
            entryId: generateUniqueId(),
            amount,
            type: 'DEBIT',
            journalId,
            accountId: debitAccountId,
            asset,
            productLine,
            timestamp
        };

        this.entries.push(debitEntry);
        this.entries.push(creditEntry);

        return { creditEntry, debitEntry };
    }

    getCurrentBalance(accountId: string, asset: string): number {
        let balance = 0;

        for (const e of this.entries) {
            if (e.accountId === accountId && e.asset === asset) {
                if (e.type === 'CREDIT') {
                    balance += e.amount;
                } 

                else if (e.type === 'DEBIT') {
                    balance -= e.amount;
                }
            }
        }

        return balance;
    }

    getTransactionHistory(accountId: string): LedgerEntry[] {
        return this.entries
            .filter(e => e.accountId === accountId)
            .sort((a, b) => a.timestamp - b.timestamp);
    }

    isBalanced(): boolean {
        let creditSum = 0;
        let debitSum = 0;

        for (const entry of this.entries) {
            if (entry.type === 'CREDIT') creditSum += entry.amount;
            else debitSum += entry.amount;
        }

        return creditSum === debitSum; 
    }
}