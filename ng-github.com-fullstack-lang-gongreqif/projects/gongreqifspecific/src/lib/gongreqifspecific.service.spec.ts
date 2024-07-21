import { TestBed } from '@angular/core/testing';

import { GongreqifspecificService } from './gongreqifspecific.service';

describe('GongreqifspecificService', () => {
  let service: GongreqifspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongreqifspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
