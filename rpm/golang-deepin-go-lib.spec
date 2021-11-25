# Run tests in check section
# disable for bootstrapping
%bcond_with check
%global import_path github.com/linuxdeepin/go-lib

%global goipath  github.com/linuxdeepin/go-lib
%global forgeurl https://github.com/linuxdeepin/go-lib

%global with_debug 1

%if 0%{?with_debug}
%global debug_package   %{nil}
%endif
%gometa

Name:           go-lib
Version:        5.7.5
Release:        1
Summary:        Go bindings for Deepin Desktop Environment development
License:        GPLv3
URL:            http://shuttle.corp.deepin.com/cache/tasks/18990/unstable-amd64/
Source0:        go-dlib_%{version}.orig.tar.xz
BuildRequires:  compiler(go-compiler)

%description
DLib is a set of Go bindings/libraries for DDE development.
Containing dbus (forking from guelfey), glib, gdkpixbuf, pulse and more.

%package devel
Summary:        %{summary}
BuildArch:      noarch
%if %{with check}
# Required for tests
BuildRequires:  deepin-gir-generator
BuildRequires:  dbus-x11
BuildRequires:  iso-codes
BuildRequires:  mobile-broadband-provider-info
BuildRequires:  golang(github.com/linuxdeepin/go-x11-client)
BuildRequires:  golang(github.com/smartystreets/goconvey/convey)
BuildRequires:  golang(gopkg.in/check.v1)
BuildRequires:  pkgconfig(gio-2.0)
BuildRequires:  pkgconfig(gdk-3.0)
BuildRequires:  pkgconfig(gdk-x11-3.0)
BuildRequires:  pkgconfig(gdk-pixbuf-xlib-2.0)
BuildRequires:  pkgconfig(libpulse)
%endif

%description devel
%{summary}.

Provides: golang(github.com/linuxdeepin/go-lib)

This package contains library source intended for
building other packages which use import path with
%{goipath} prefix.

%prep
%setup -q -n  go-dlib-%{version}
%forgeautosetup -n go-dlib-%{version}

%install
install -d -p %{buildroot}/%{gopath}/src/%{import_path}/
for file in $(find . -iname "*.go" -o -iname "*.c" -o -iname "*.h") ; do
    install -d -p %{buildroot}/%{gopath}/src/%{import_path}/$(dirname $file)
    cp -pav $file %{buildroot}/%{gopath}/src/%{import_path}/$file
    echo "%%{gopath}/src/%%{import_path}/$file" >> devel.file-list
done

cp -pav README.md %{buildroot}/%{gopath}/src/%{goipath}/README.md
cp -pav CHANGELOG.md %{buildroot}/%{gopath}/src/%{goipath}/CHANGELOG.md
echo "%%{gopath}/src/%%{goipath}/README.md" >> devel.file-list
echo "%%{gopath}/src/%%{goipath}/CHANGELOG.md" >> devel.file-list

%if %{with check}
%check
%gochecks
%endif

%files devel -f devel.file-list
%doc README.md
%license LICENSE

%changelog
* Wed Mar 12 2021 uoser <uoser@uniontech.com> - 5.7.5-1
- Update to 5.7.5
