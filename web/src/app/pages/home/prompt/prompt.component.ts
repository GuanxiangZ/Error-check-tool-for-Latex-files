import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-prompt',
  templateUrl: './prompt.component.html',
  styleUrls: ['./prompt.component.scss']
})
export class PromptComponent implements OnInit {

  @Input()
  prompt = '';

  @Input()
  statement = '';

  constructor() { }

  ngOnInit(): void {
  }

}
