import { NgModule } from '@angular/core';
import {CommonModule} from '@angular/common';
import { HomeRoutingModule } from './home-routing.module';

import {HomeComponent} from './home.component';
import {NzUploadModule} from 'ng-zorro-antd/upload';
import {NzIconModule} from 'ng-zorro-antd/icon';
import {NzButtonModule} from 'ng-zorro-antd/button';
import {NzMessageService} from 'ng-zorro-antd/message';
import {NzModalModule} from 'ng-zorro-antd/modal';
import {FormsModule} from '@angular/forms';
import {NzSpaceModule} from 'ng-zorro-antd/space';
import {NzCardModule} from 'ng-zorro-antd/card';
import {NzInputModule} from 'ng-zorro-antd/input';
import {NzSkeletonModule} from 'ng-zorro-antd/skeleton';
import {NzCollapseModule} from 'ng-zorro-antd/collapse';
import { PromptComponent } from './prompt/prompt.component';
import {NzSelectModule} from 'ng-zorro-antd/select';
import {NzToolTipModule} from 'ng-zorro-antd/tooltip';


@NgModule({
  imports: [
    CommonModule,
    HomeRoutingModule,
    NzUploadModule,
    NzIconModule,
    NzButtonModule,
    NzModalModule,
    FormsModule,
    NzSpaceModule,
    NzCardModule,
    NzInputModule,
    NzSkeletonModule,
    NzCollapseModule,
    NzSelectModule,
    NzToolTipModule,
  ],
  declarations: [HomeComponent, PromptComponent],
  exports: [HomeComponent],
  providers: [
    NzMessageService,
  ]
})
export class HomeModule { }
