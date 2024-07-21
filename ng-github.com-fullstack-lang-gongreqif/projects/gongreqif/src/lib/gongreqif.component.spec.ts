import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongreqifComponent } from './gongreqif.component';

describe('GongreqifComponent', () => {
  let component: GongreqifComponent;
  let fixture: ComponentFixture<GongreqifComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongreqifComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongreqifComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
