export interface Report {
  lines: Line[];
  errorLineIndexes: number[];
  title: string;
}

export interface Line {
  matches: Match[];
  lineIndex: number;
  len: number;
  segments: Segment[];
  start: number;
}

export interface Segment {
  text: string;
  type: string;
  id: string;
}

export interface Match {
  id: string;
  reason: string;
  matchPreString: string;
  matchPostString: string;
  globalErrorStartPosition: number;
  globalErrorEndPosition: number;
  errorString: string;
  replacements: Replacement[];
  issueType: string;
  issueDescription: string;
}

export interface Replacement {
  index: number;
  replacementString: string;
}
