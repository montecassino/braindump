import { test, suite } from 'node:test';
import assert from 'node:assert/strict';

const CASH_ACCOUNT = 'CASH_MAIN_USD'
const ACCOUNT_A = 'USER_A_EQUITY'
const ACCOUNT_B = 'USER_B_CRYPTO'
const BTC = 'BTC'
const USD = 'USD'

suite('Double entry ledger test suite', () => {
    let ledger : Ledger

    test.beforeEach(() => {
        ledger = new Ledger()
    })

    test('should have correct initial state', () => {
        assert.strictEqual(ledger.entries.length, 0, "Ledger length should be 0");
        assert.strictEqual(ledger.getCurrentBalance(ACCOUNT_A, USD), 0, "Account A balance should be 0");
        assert.strictEqual(ledger.isBalanced(), true, "Ledger should be balanced initially");
    })

    test('should correctly record a simple USD deposit', () => {
        ledger.recordTransaction(CASH_ACCOUNT, ACCOUNT_A, 69, USD, 'EQUITY')

        assert.strictEqual(ledger.entries.length, 2, "Credit and debit entries must be added to the ledger");
        assert.strictEqual(ledger.isBalanced(), true, "Ledger must be balanced");
        assert.strictEqual(ledger.getCurrentBalance(CASH_ACCOUNT, USD), -69, "CASH_ACCOUNT balance should be -69 (debit)");
        assert.strictEqual(ledger.getCurrentBalance(ACCOUNT_A, USD), 69, "ACCOUNT_A balance should be 69 (credit)");
    });

    test('should handle cross-asset transactions', () => {
        ledger.recordTransaction(CASH_ACCOUNT, ACCOUNT_A, 100, USD)
        ledger.recordTransaction(ACCOUNT_B, CASH_ACCOUNT, 1, BTC)

        asset.strictEqual(ledger.entries.length, 4, "4 entries should've been added")
        assert.strictEqual(ledger.isBalanced(), true, "Ledger must be balanced")

        assert.strictEqual(ledger.getCurrentBalance(CASH_ACCOUNT, BTC), 1, 'Credited 1 BTC')
        assert.strictEqual(ledger.getCurrentBalance(ACCOUNT_B), -1, 'Debited 1 BTC')

        assert.strictEqual(ledger.getCurrentBalance(ACCOUNT_A, USD), 100, "ACCOUNT_A USD balance should be 100");
    })
})