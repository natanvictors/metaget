# 2026-11-26

## omit locale in LoLalytics URLs
Considered: en_us (doesn't exist), pt_br (pinned), omit locale segment
Chose: omit locale
Why: verified, omitting returns English, which is the intended parsing language
Revisit if: default locale changes, or a locale-specific number appears

## data grabbing technique
Considered: raw text parsing, structural selectors
Chose: structural selectors
Why: the same phrasing can be used to different data in specific contexts
Revisit if: big layout/UI change on the website