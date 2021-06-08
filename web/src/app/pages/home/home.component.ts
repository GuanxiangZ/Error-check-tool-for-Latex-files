import {
  AfterViewInit, ChangeDetectorRef,
  Component,
  ElementRef,
  HostListener,
  OnInit,
  ViewChild, ViewEncapsulation,
} from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import {NzUploadChangeParam, NzUploadComponent, NzUploadFile} from 'ng-zorro-antd/upload';
import {HttpClient} from '@angular/common/http';

import {Observable} from 'rxjs';
import {Line, Match, Replacement, Report, Segment} from '../report';
import {NzSkeletonParagraph} from 'ng-zorro-antd/skeleton';
import {saveAs} from 'file-saver';


function getFileContent(file: File): Promise<string> {
  return new Promise((resolve, reject) => {
    if (!file) {
      resolve('');
    }
    const reader = new FileReader();
    reader.onloadend = (e) => {
      const text = reader.result?.toString();
      resolve(text ?? '');
    };

    reader.readAsText(file);
  });
}



@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class HomeComponent implements OnInit, AfterViewInit {
  baseUrl = 'api/upload';

  supportLanguages: Map<string, string> = new Map<string, string>();
  supportLanguagesName: string[] = [
    'English (US)',
    'English (British)',
    'English (Australia)',
    'English (Canada)',
    'English (New Zealand)',
    'English (South Africa)',
    'French',
    'German (Austria)',
    'German (Germany)',
    'German (Switzerland)',
    'Greek',
    'Irish'
  ];

  constructor(private http: HttpClient, private msg: NzMessageService, private changeDetectorRef: ChangeDetectorRef) {}

  response: Response = {} as Response;

  report: Report  = {} as Report;

  @ViewChild('uploadComponent', {static: false})
  uploadComponent: NzUploadComponent | undefined;

  @ViewChild('mainContainer', { static: false })
  mainContainer: ElementRef | undefined;

  @ViewChild('previewContainer', { static: false })
  previewContainer: ElementRef | undefined;

  @ViewChild('textarea', { static: false })
  textarea: ElementRef<HTMLTextAreaElement> | undefined;

  currentFocusError = '';
  currentFocusErrorLine = -1;
  currentLanguage = 'English (US)';

  previewAreaWidth = 0;

  textAreaMaxHeight = 0;

  textAreaLineHeight = 20;
  textAreaRows = 0;

  uploading = false;

  file: NzUploadFile | undefined;

  fileName: string | undefined = '';

  fileContent = '';

  lines: string[] = [];

  errorLineMap: Map<number, Line> = new Map<number, Line>();

  correctedLineMap: Map<number, string> = new Map<number, string>();

  beforeUpload = (file: NzUploadFile): boolean | Observable<boolean> =>  {

    this.fileName = file.name;

    this.uploading = true;
    const suffix = this.fileName.substring(this.fileName.lastIndexOf('.'));

    if (suffix  === '.tex' || suffix === '.latex' || file.type === 'application/x-tex') {
      if (this.uploadComponent && this.uploadComponent?.nzFileList.length !== 0) {
        this.uploadComponent.nzFileList = [];
      }
      return true;
    }

    this.msg.error(`File type error: please upload a file with .tex or .latex suffix`);
    this.clearFile();
    this.uploading = false;

    return false;
  }

  clearFile = () => {
    this.fileContent = '';
    this.fileName = '';
    this.lines = [];
    this.file = undefined;
    this.report = {} as Report;
    this.errorLineMap = new Map<number, Line>();
    this.correctedLineMap = new Map<number, string>();
  }

  handleChange = async (info: NzUploadChangeParam) => {

    if (info.file.status === 'removed') {
      this.clearFile();
      this.changeDetectorRef.markForCheck();
      return;
    }

    if (info.file.status === 'done') {
      if (info.file.originFileObj !== undefined) {
        this.fileContent = await getFileContent(info.file.originFileObj);

        this.textAreaMaxHeight = this.mainContainer?.nativeElement.offsetHeight;

        this.lines = this.fileContent.split('\n');

        this.fileName = info.file.originFileObj.name;
      }

      this.msg.success(`${info.file.name} file uploaded successfully`);

      this.response = info.file.response;
      if (this.response.err !== 0) {
        const specError = 'rpc error: code = Unknown desc = Exception calling application:';
        if (this.response.err === 2) {
          let msg = this.response.msg;
          const pos = msg.indexOf(specError);
          msg = msg.substr(pos + specError.length);
          this.msg.error(`TeX File Syntax Error: ${msg}`, {
            nzDuration: 5000,
          });
          this.fileContent = '';
        } else {
          this.msg.error(`Server error: ${this.response.msg}`, {
            nzDuration: 5000,
          });

          this.fileContent = '';
        }
        this.uploading = false;
        return;
      }

      this.report = this.response.report;
      if (this.report.lines.length === 0) {
        this.msg.success(`${info.file.name} file does not have grammatical errors!`, {
          nzDuration: 5000,
        });
        this.uploading = false;
        this.fileContent = '';
        return;
      }

      for (const line of this.report.lines) {
        this.errorLineMap.set(line.lineIndex, line);
      }
      this.fileContent = '';
      this.uploading = false;
    } else if (info.file.status === 'error') {
      this.uploading = false;
      this.msg.error(`${info.file.name} file upload failed.`, {
        nzDuration: 5000,
      });
    }
  }


  @HostListener('window:resize', ['$event'])
  onResize(event: Event): void {
    if (this.mainContainer !== undefined) {
      this.textAreaMaxHeight = this.mainContainer.nativeElement.offsetHeight;
      this.textAreaRows = Math.floor(this.textAreaMaxHeight / this.textAreaLineHeight);
      this.previewAreaWidth = this.mainContainer?.nativeElement.offsetWidth / 2;
    }
  }

  ngOnInit(): void {
    this.supportLanguages.set('English (US)', 'en-US');
    this.supportLanguages.set('English (British)', 'en-GB');
    this.supportLanguages.set('English (Australia)', 'en-AU');
    this.supportLanguages.set('English (Canada)', 'en-CA');
    this.supportLanguages.set('English (New Zealand)', 'en-NZ');
    this.supportLanguages.set('English (South Africa)', 'en-ZA');
    this.supportLanguages.set('French', 'fr');
    this.supportLanguages.set('German (Austria)', 'de-AT');
    this.supportLanguages.set('German (Germany)', 'de-DE');
    this.supportLanguages.set('German (Switzerland)', 'de-CH');
    this.supportLanguages.set('Greek', 'el-GR');
    this.supportLanguages.set('Irish', 'ga-IE');
  }

  ngAfterViewInit(): void {

    this.textAreaMaxHeight = this.mainContainer?.nativeElement.offsetHeight;
    this.textAreaRows = Math.floor(this.textAreaMaxHeight / this.textAreaLineHeight);

    this.previewAreaWidth = this.mainContainer?.nativeElement.offsetWidth / 2;

    if (this.textarea !== undefined) {
      this.textarea.nativeElement.style.height = this.textAreaMaxHeight.toString();
    }
    this.changeDetectorRef.detectChanges();
  }


  getSkeletonRows(height: number): NzSkeletonParagraph {
    return {
      rows: Math.floor(height / 40),
    };
  }

  getHeader(lineIndex: number): string {
    return 'Line number: ' + (lineIndex + 1).toString();
  }

  openCollapse = (line: Line) => {
    const el = document.getElementById('line_number_' + line.lineIndex.toString());
    el?.scrollIntoView({behavior: 'smooth'});
    if (line.matches.length > 0) {
      this.currentFocusError = line.matches[0].id;
    }
    this.currentFocusErrorLine = line.lineIndex;
    this.changeDetectorRef.markForCheck();
  }

  openErrorCollapse = (event: Event, err: string) => {
    event.stopPropagation();
    this.currentFocusError = err;
    this.changeDetectorRef.markForCheck();
  }

  isActive(lineIndex: number): boolean {
    return lineIndex === this.currentFocusErrorLine;
  }

  getOtherSuggestions(replacements: Replacement[]): Replacement[]  {
    if (replacements.length > 1) {
      return replacements.slice(1);
    }
    return [];
  }

  getStyle(seg: Segment): string {
    if (seg.type === 'normal') {
      return 'normal-segment';
    }
    return 'error-segment';
  }

  getLine(idx: number): boolean {
    if (!this.report.errorLineIndexes) {
      return true;
    }
    return !(this.report.errorLineIndexes.includes(idx));
  }

  getLineId(idx: number): string {
    return 'line_number_' + idx.toString();
  }

  getSegments(idx: number): Segment[] {
    for (const line of this.report.lines) {
      if (line.lineIndex === idx) {
        return line.segments;
      }
    }
    return [];
  }

  getSkeletonParagraph(height: number): NzSkeletonParagraph {
    return {
      rows: Math.floor(height / 40),
    };
  }

  clickError = async (lineIndex: number, id: string): Promise<void> =>  {
    this.currentFocusError = id;
    this.currentFocusErrorLine = lineIndex;

    this.changeDetectorRef.markForCheck();
    return new Promise<void>((resolve, reject) => {

      setTimeout(() => {
        const pel = document.getElementById(this.getReportLineID(lineIndex));
        pel?.scrollIntoView({behavior: 'smooth'});

        setTimeout(() => {
          const el = document.getElementById(this.getCollapseErrorID(id));
          el?.scrollIntoView({behavior: 'smooth'});
        }, 50);

      }, 300);
    });
  }

  getCurrentLanguageCode = (): string =>  {
    return this.supportLanguages.get(this.currentLanguage) ?? 'en-US';
  }

  getReportLineID(lineIndex: number): string {
    return 'report_'  + lineIndex.toString();
  }

  getCollapseErrorID(id: string): string {
    return 'report_collapse_' + id;
  }

  downloadFile = (_: NzUploadFile): void => {
    const content: string[] = [];

    for (let i = 0; i < this.lines.length; i++) {
      const line = this.correctedLineMap.get(i);
      if (line) {
        content.push(line + '\n');
      } else {
        content.push(this.lines[i] + '\n');
      }
    }

    const prefix =  this.fileName?.substring(0, this.fileName?.lastIndexOf('.'));

    const file = new File(content, `${prefix}-corrected.tex`, {type: 'application/x-tex'});
    saveAs(file);
  }

  fixError = (lineIndex: number, replIdx: number, match: Match) => {
    const el = document.getElementById(match.id);
    if (el !== null) {
      el.innerText = match.replacements[replIdx].replacementString;
      this.removeError(el, lineIndex, match);
    }
  }

  dismissError = ($event: MouseEvent, lineIndex: number, match: Match) => {
    $event.stopPropagation();
    const el = document.getElementById(match.id);
    if (el !== null) {
      this.removeError(el, lineIndex, match);
    }
  }

  removeError = (el: HTMLElement, lineIndex: number, match: Match)  => {
    el.classList.remove('error-segment');
    el.classList.add('normal-segment');

    const pel = el.parentElement;
    let lineContent = '';
    if (pel !== null) {

      // tslint:disable-next-line:prefer-for-of
      for (let i = 0; i < pel.children.length; i++) {
        const cld = pel.children[i] as HTMLElement;
        lineContent += cld.innerText.trim() + ' ';
      }

      this.correctedLineMap.set(lineIndex, lineContent);

      const line: Line = this.errorLineMap.get(lineIndex)!;

      if (line.matches.length === 1) {
        const lineEL = document.getElementById(this.getReportLineID(lineIndex));
        if (lineEL) {
          lineEL.hidden = true;
        }
      }

      let idx = -1;

      const matchLen =  line.matches.length;
      // tslint:disable-next-line:prefer-for-of
      for (let i = 0; i < matchLen; i++) {
        if (line.matches[i].id === match.id) {
          idx = 1;
          break;
        }
      }

      if (idx !== -1) {
        const newMatches = [];
        for (const m of line.matches) {
          if (m.id !== match.id) {
            newMatches.push(m);
          }
        }
        line.matches = newMatches;
      }
    }

    if (this.errorLineMap.size === 0) {
      this.msg.success(`All errors in ${this.fileName} file has been corrected!`, {
        nzDuration: 5000,
      });
    }
  }
}


interface Response {
  msg: string;
  err: number;
  report: Report;
}







