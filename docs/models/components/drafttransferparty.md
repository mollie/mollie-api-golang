# DraftTransferParty

A party involved in the draft transfer, representing either the debtor (sender) or creditor
(recipient). Contains the party's name and account details.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `FullName`                                                                                   | `string`                                                                                     | :heavy_check_mark:                                                                           | The full name of the account holder.                                                         | Jan Jansen                                                                                   |
| `Account`                                                                                    | [components.DraftTransferPartyAccount](../../models/components/drafttransferpartyaccount.md) | :heavy_check_mark:                                                                           | The bank account details of the party.                                                       |                                                                                              |