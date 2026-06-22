#!/usr/bin/env python3
import argparse
import re
import sys


def main() -> int:
    parser = argparse.ArgumentParser(
        description="""
            Ensure changes in source files are accompanied by changes in destination files using regex
            and logical operators.
            """,
    )

    parser.add_argument(
        "--source",
        dest="sources",
        action="append",
        required=True,
        help="Regex pattern for source files",
    )
    parser.add_argument(
        "--destination",
        dest="destinations",
        action="append",
        required=True,
        help="Regex pattern for destination files",
    )

    # Add flags for logical operators, defaulting to OR to match previous behavior
    parser.add_argument(
        "--source-logic",
        choices=["AND", "OR"],
        default="OR",
        help="Logic operator for sources (default: OR)",
    )
    parser.add_argument(
        "--destination-logic",
        choices=["AND", "OR"],
        default="OR",
        help="Logic operator for destinations (default: OR)",
    )

    parser.add_argument("filenames", nargs="*", help="List of staged filenames")

    args = parser.parse_args()

    try:
        source_patterns = [re.compile(p) for p in args.sources]
        dest_patterns = [re.compile(p) for p in args.destinations]
    except re.error as e:
        print(f"❌ ERROR: Invalid regular expression provided. Details: {e}")
        return 1

    # Use sets to keep track of the indices of the patterns that were matched
    matched_sources = set()
    matched_destinations = set()

    for filename in args.filenames:
        # Check against all source patterns
        for idx, src_pattern in enumerate(source_patterns):
            if src_pattern.search(filename):
                matched_sources.add(idx)

        # Check against all destination patterns
        for idx, dest_pattern in enumerate(dest_patterns):
            if dest_pattern.search(filename):
                matched_destinations.add(idx)

    # Evaluate Source Logic
    if args.source_logic == "OR":
        sources_triggered = len(matched_sources) > 0
    else:  # AND
        sources_triggered = len(matched_sources) == len(source_patterns)

    # Evaluate Destination Logic
    if args.destination_logic == "OR":
        destinations_satisfied = len(matched_destinations) > 0
    else:  # AND
        destinations_satisfied = len(matched_destinations) == len(dest_patterns)

    # If the source condition is met, but the destination condition is NOT met, fail the hook.
    if sources_triggered and not destinations_satisfied:
        print("❌ ERROR: Required synchronized changes missing.")
        print(f"Source logic ({args.source_logic}) was triggered for patterns: {args.sources}")
        print(f"But destination logic ({args.destination_logic}) was NOT satisfied for patterns: {args.destinations}")
        return 1

    return 0


if __name__ == "__main__":
    sys.exit(main())
