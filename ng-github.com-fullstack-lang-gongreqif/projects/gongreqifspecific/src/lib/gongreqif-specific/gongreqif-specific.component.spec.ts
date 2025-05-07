import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongreqifSpecificComponent } from './gongreqif-specific.component';

describe('GongreqifSpecificComponent', () => {
  let component: GongreqifSpecificComponent;
  let fixture: ComponentFixture<GongreqifSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongreqifSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongreqifSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
