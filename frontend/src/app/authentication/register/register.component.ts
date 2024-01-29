import {Component, DestroyRef, inject, OnInit} from '@angular/core'
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms"
import {Router, RouterModule} from "@angular/router"
import {AuthenticationService} from "../../service/authentication.service"
import {takeUntilDestroyed} from "@angular/core/rxjs-interop"
import {RegisterDetail} from "../../model/register-detail";

interface RegisterForm {
  firstName: FormControl<string>
  lastName: FormControl<string>
  phoneNumber: FormControl<string>
  email: FormControl<string>
  password: FormControl<string>
  confirmPassword: FormControl<string>
}

@Component({
  selector: 'register-component',
  standalone: true,
  imports: [ReactiveFormsModule, RouterModule],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export default class RegisterComponent implements OnInit {
  protected registerForm: FormGroup<RegisterForm>
  protected isSubmitting: boolean

  private destroyRef = inject(DestroyRef)

  constructor(
    private readonly router: Router,
    private readonly authenticationService: AuthenticationService,
  ) {
    this.registerForm = new FormGroup<RegisterForm>({
      firstName: new FormControl("", {
        validators: [Validators.required],
        nonNullable: true,
      }),
      lastName: new FormControl("", {
        validators: [Validators.required],
        nonNullable: true,
      }),
      phoneNumber: new FormControl("", {
        validators: [Validators.required],
        nonNullable: true,
      }),
      email: new FormControl("", {
        validators: [Validators.required],
        nonNullable: true,
      }),
      password: new FormControl("", {
        validators: [Validators.required],
        nonNullable: true,
      }),
      confirmPassword: new FormControl("", {
        validators: [Validators.required],
        nonNullable: true,
      }),
    })
    this.isSubmitting = false
  }

  ngOnInit(): void {
  }

  submitForm(): void {
    this.isSubmitting = true

    let observable = this.authenticationService
      .register(this.registerForm.value as RegisterDetail)

    observable.pipe(takeUntilDestroyed(this.destroyRef)).subscribe({
      next: (response) => {
        if (response !== null && response.success) {
          this.router.navigate(["/"])
        }
        this.isSubmitting = false
      },
      error: (err) => {
        console.log(err)
        this.isSubmitting = false
      },
    })
  }
}
