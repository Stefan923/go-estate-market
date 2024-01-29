import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListPropertiesComponent } from './list-properties.component';

describe('ListPropertiesComponent', () => {
  let component: ListPropertiesComponent;
  let fixture: ComponentFixture<ListPropertiesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ListPropertiesComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ListPropertiesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
