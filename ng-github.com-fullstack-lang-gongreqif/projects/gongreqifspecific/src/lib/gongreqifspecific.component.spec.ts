import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongreqifspecificComponent } from './gongreqifspecific.component';

describe('GongreqifspecificComponent', () => {
  let component: GongreqifspecificComponent;
  let fixture: ComponentFixture<GongreqifspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongreqifspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongreqifspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
