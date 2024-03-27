# Superform Task (Backend)

# Task 1: Superform CLI Analyzer Tool

This tool is designed to analyze a set of cross-chain withdrawal transactions, specifically focusing on calculating the average processing time. It uses data from successfully completed transactions obtained via the Superscan API, considering transactions between specified source and destination chain IDs, and involving a combination of two AMBs (LayerZero, Hyperlane, and Wormhole).

## Features

- Filter transactions based on source and destination chain IDs and AMBs used and fetch all the missing data needed for further optimizations (`outputAmount` after the `UPDATE_HASH` goes through, etc.).
- Calculate the processing time for each transaction as the difference between the `delayed_timestamp` (after the `UPDATE_HASH` goes through) and the `initial_timestamp`.
- Compute the average processing time across all transactions that meet the filtering criteria.
- Compute the Time Adjustment Factor (TAF) needed for the Task 2.

## Requirements

- Go programming language (version 1.15 or higher recommended)

## Installation

Before installing and running this tool, ensure you have Go installed on your system. You can download Go from [the official website](https://golang.org/dl/).

To build and install the tool, follow these steps:

1. Clone the repository or download the source code to your local machine.
2. Navigate to the directory containing the source code.
3. Run the build script: `./build.sh`
4. If you wish to have the tool available system-wide, move the binary to a location in your system's PATH, e.g., `/usr/local/bin`.

## Usage

To use the tool, execute it from the command line with the required parameters. Here are the available flags:

- `sourceChainID`: The ID of the source blockchain network (default: 42161).
- `destinationChainID`: The ID of the destination blockchain network (default: 10).
- `ambs`: A comma-separated list of two AMBs to use, chosen from Layerzero, Hyperlane, and Wormhole.
- `output`: Optional. Specifies the path to a JSON output file where the processed transactions will be saved.

## Notes

- The tool has a cap of maximum 2000 transactions to fetch from the Superscan API without filtering. That can be removed once the retries on blockchain requests are implemented, custom RPCs, etc. so all transactions can be processed.
  - Also, the tool can be optimized to fetch only the transactions that are needed for calculations by filtering them out on the API level.
- The tool can be optimized further by implementing concurrent processing of transactions, error handling, and retries for failed requests.
- Tests can be added to ensure the correctness of the processing logic and the accuracy of the calculated metrics.

### Example Command

```shell
./bin/superformcli -sourceChainID=42161 -destinationChainID=10 -ambs="Layerzero,Hyperlane" -output="transactions.json"
```

The output will be displayed in the terminal and processed transactions will be saved to the specified output file in JSON format.

Example output for the above command:

```shell
source Chain ID: 42161
destination Chain ID: 10
AMBs: [Layerzero Hyperlane]
processed transactions will be written to: transactions.json
fetching maximum 2000 transactions for demonstration purposes...
processing 2000 transactions...
processing transaction: 0x53985cacee4538252a72de986923b695e070fc0131b59578ef9161356eb51e2e
processing transaction: 0x1c80e0802b823ac16299548207dba60132e1c7952394d22b2ea35ec03f72644d
processing transaction: 0x254ec178fe3d1f0f83449e5444a2950861195aeafecc65d85050f2aa80b07834
processing transaction: 0x15a677d8e05b2fd2b0845f2210c480bd552120c5559e94c7118890b1a48921fe
processing transaction: 0x95208e4face18fca2d16505b130287c45157d5d32dad2161074df2508afae382
processing transaction: 0xd9c0379b4f16bf512768296a0416b7dc055a0519b364d28d19004cfcad9c0fc0
processing transaction: 0xf2757d3047c3b5568f6f0535ecbafcabd9dd2dd7ac83a086d7345df0fadb3271

...

processing transaction: 0xe6676cfa820e22060a3411af70f68a62238695b0e96474f0b5505693da2280d2
processing transaction: 0xca4725adeb2f4ea8b7567abbe6c09e4616d02666e7de53b44beffac488106f53
processing transaction: 0xb2014a0da5b23929a6f4dca4899a1ee911ab21480ae963f8c365642ab8f66f28
processing transaction: 0x2ef327bae16758094e9f5b55c9a2080dc9ca480d5dcfa1d93fc00b6ed403d928

average time until the transaction can be processed is 2m38.69090909s and the time adjustment factor: -5.79%
```

The contents of the `transactions.json` file will contain the processed transactions in JSON format. For example:

```json
[
  {
    "type": "withdrawal",
    "id": 5428,
    "hash": "0x1c80e0802b823ac16299548207dba60132e1c7952394d22b2ea35ec03f72644d",
    "chain": "42161",
    "status": "confirmed",
    "superform_id": "62771017356061639854904882663925984399197993391790181127465",
    "amount": "34543341",
    "initial_timestamp": 1711490465,
    "delayed_timestamp": 1711490557,
    "initial_output_amount": "35678282",
    "delayed_output_amount": "33444488"
  },
  {
    "type": "withdrawal",
    "id": 5425,
    "hash": "0x254ec178fe3d1f0f83449e5444a2950861195aeafecc65d85050f2aa80b07834",
    "chain": "42161",
    "status": "confirmed",
    "superform_id": "62771017356061639854904882663925984399197993391790181127465",
    "amount": "12583986",
    "initial_timestamp": 1711488095,
    "delayed_timestamp": 1711488185,
    "initial_output_amount": "12997334",
    "delayed_output_amount": "12183778"
  },
  {
    "type": "withdrawal",
    "id": 5408,
    "hash": "0x15a677d8e05b2fd2b0845f2210c480bd552120c5559e94c7118890b1a48921fe",
    "chain": "42161",
    "status": "processing",
    "superform_id": "62771017356061639854904882663925984399197993391790181127465",
    "amount": "97282723",
    "initial_timestamp": 1711480155,
    "delayed_timestamp": 1711480263,
    "initial_output_amount": "100475450",
    "delayed_output_amount": "94191399"
  },
  {
    "type": "withdrawal",
    "id": 5387,
    "hash": "0x38566c169d68709dbec00669ba536fc2083346ed57016e0990f7db4489731092",
    "chain": "42161",
    "status": "confirmed",
    "superform_id": "62771017356061639854904882663925984399197993391790181127465",
    "amount": "263278876",
    "initial_timestamp": 1711468403,
    "delayed_timestamp": 1711468507,
    "initial_output_amount": "271908507",
    "delayed_output_amount": "254923009"
  },
  {
    "type": "withdrawal",
    "id": 5380,
    "hash": "0x1598fced62f192412026659ca8895d29858475d1aa8eeab7f52f62061b415e74",
    "chain": "42161",
    "status": "confirmed",
    "superform_id": "62771017355823347281171697645255832504345359639728524163157",
    "amount": "51280334919394445",
    "initial_timestamp": 1711465279,
    "delayed_timestamp": 1711465385,
    "initial_output_amount": "52076314268914991",
    "delayed_output_amount": "50496520633495298"
  },
  {
    "type": "withdrawal",
    "id": 5374,
    "hash": "0x6d4f91e72855a54bb59fd15409afc202f062f1dfb28b6f0fd2ca618125d31c3a",
    "chain": "42161",
    "status": "confirmed",
    "superform_id": "62771017356061639854904882663925984399197993391790181127465",
    "amount": "506232754",
    "initial_timestamp": 1711463842,
    "delayed_timestamp": 1711463937,
    "initial_output_amount": "522817621",
    "delayed_output_amount": "490173746"
  }
]
```

These transactions are crucial for the Task 2 implementation since they contain the necessary data for the Time Adjustment Factor (TAF) calculation:

- `initial_output_amount`: The initial output amount at the time of transaction initiation.
- `delayed_output_amount`: The output amount after the `UPDATE_HASH` goes through.
- `initial_timestamp`: The timestamp when the transaction was initiated.
- `delayed_timestamp`: The timestamp when the transaction can be processed.

## Task 2 - Withdrawal Output Amount Adaptation Methodology

This outlines the methodology used to adapt the estimated output amount for withdrawals in the Superform protocol, taking into consideration the processing time delays and their impact on the final output amount. The approach utilizes two key formulas: the Time Adjustment Factor (TAF) and the Adapted Output Amount (AAO).

`NOTE`: the formula can be further extended to include fees, gas costs, and other factors that may impact the final output amount.

## Time Adjustment Factor (TAF)

The Time Adjustment Factor (TAF) is a metric that reflects the average percentage change in transaction output amounts relative to the initial output amount, weighted by the processing time for each transaction. This factor helps in understanding how processing delays can affect the final output amounts.

### Formula

The TAF is calculated as follows:

```TAF = (Σ((DAO - IAO) / IAO) * TD) / ΣTD * 100```

Where:
- `DAO` is the Delayed Output Amount at the time the transaction is processed.
- `IAO` is the Initial Output Amount at the time of transaction initiation.
- `TD` is the Time Delay, i.e., the difference between the delayed timestamp and the creation timestamp.
- The formula calculates a time-weighted average of the percentage changes in output amount, taking into account the duration of each transaction's delay time.

## Adapted Output Amount (AAO)

The Adapted Output Amount (AAO) uses the TAF to adjust the initial output amount estimates, providing a more accurate prediction of the expected output amount at the time of transaction processing.

### Formula

The AAO is calculated using the initial output amount and the TAF as follows:

```AAO = IAO + (IAO * TAF / 100)```

Where:
- `AAO` is the Adapted Output Amount, reflecting the adjusted estimate after considering the TAF.
- `IAO` is the Initial Output Amount, the original estimate provided at the time of transaction initiation.
- `TAF` is the Time Adjustment Factor, calculated as described above.

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.