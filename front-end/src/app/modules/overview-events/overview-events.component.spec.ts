import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OverviewEventsComponent } from './overview-events.component';

describe('OverviewEventsComponent', () => {
  let component: OverviewEventsComponent;
  let fixture: ComponentFixture<OverviewEventsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OverviewEventsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OverviewEventsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
